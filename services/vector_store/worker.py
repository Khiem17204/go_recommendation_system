# worker.py
import threading
import queue
import logging
from db import fetch_unembedded_decks, fetch_deck_cards, insert_deck_embedding
from sentence_transformers import SentenceTransformer

logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

model = SentenceTransformer("all-MiniLM-L6-v2")
deck_queue = queue.Queue()

NUM_WORKERS = 10

def worker():
    while True:
        deck = deck_queue.get()
        if deck is None:
            break
        try:
            cards = fetch_deck_cards(deck['id'])
            text = " ".join(cards)
            embedding = model.encode([text])[0]  # single embedding
            insert_deck_embedding(deck['id'], deck['deck_name'], embedding)
            logging.info(f"✅ Embedded deck {deck['id']}: {deck['deck_name']}")
        except Exception as e:
            logging.error(f"❌ Failed to embed deck {deck['id']}: {e}")
        finally:
            deck_queue.task_done()

def main():
    decks = fetch_unembedded_decks()
    logging.info(f"Found {len(decks)} unembedded decks to process.")

    for deck in decks:
        deck_queue.put(deck)

    threads = []
    for _ in range(NUM_WORKERS):
        t = threading.Thread(target=worker)
        t.start()
        threads.append(t)

    deck_queue.join()

    for _ in range(NUM_WORKERS):
        deck_queue.put(None)
    for t in threads:
        t.join()

    logging.info("All deck embeddings completed.")

if __name__ == '__main__':
    main()
