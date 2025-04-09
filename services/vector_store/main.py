# vectorstore/service.py
from concurrent import futures
import grpc
import vector_pb2
import vector_pb2_grpc
from store import CardStore, DeckStore
import sys

class BaseVectorStore:
    def search(self, cards):
        raise NotImplementedError

class CardVectorStore(BaseVectorStore):
    def __init__(self):
        self.store = CardStore()

    def search(self, cards):
        grouped = [cards[i:i+3] for i in range(0, len(cards), 3)]
        freq = {}
        for group in grouped:
            result = self.store.search_similar_cards(list(group))
            for card in result:
                freq[card] = freq.get(card, 0) + 1
        top = sorted(freq.items(), key=lambda x: x[1], reverse=True)[:5]
        return [c for c, _ in top]

class DeckVectorStore(BaseVectorStore):
    def __init__(self):
        self.store = DeckStore()

    def search(self, cards):
        combined = " ".join(cards)
        return self.store.search_similar_decks(combined)

class VectorStoreServicer(vector_pb2_grpc.VectorServiceServicer):
    def __init__(self, store: BaseVectorStore):
        self.store = store

    def SearchSimilarCards(self, request, context):
        result = self.store.search(request.cards)
        return vector_pb2.CardResult(cards=result)

    def SearchSimilarDecks(self, request, context):
        result = self.store.search(request.cards)
        return vector_pb2.DeckResult(decks=result)


def serve(mode):
    if mode == "card":
        vector = CardVectorStore()
    elif mode == "deck":
        vector = DeckVectorStore()
    else:
        raise ValueError("Invalid mode: choose 'card' or 'deck'")

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    vector_pb2_grpc.add_VectorServiceServicer_to_server(VectorStoreServicer(vector), server)
    server.add_insecure_port('[::]:60051')
    server.start()
    print(f"VectorStore gRPC service ({mode}) running on port 60051")
    server.wait_for_termination()

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python service.py [card|deck]")
        sys.exit(1)
    serve(sys.argv[1])