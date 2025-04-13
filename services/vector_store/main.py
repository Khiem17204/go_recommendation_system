# vectorstore/faiss_service.py
from concurrent import futures
import grpc
import numpy as np
import faiss
import os
import sys
from sentence_transformers import SentenceTransformer

import db
import search_pb2
import search_pb2_grpc

INDEX_PATH = os.getenv("INDEX_PATH", "deck_faiss.index")
ID_MAP_PATH = os.getenv("ID_MAP_PATH", "deck_ids.npy")

class FaissVectorStore:
    def __init__(self):
        self.model = SentenceTransformer("all-MiniLM-L6-v2")
        self.index = None
        self.deck_ids = []
        self._load_embeddings()

    def _load_embeddings(self):
        if os.path.exists(INDEX_PATH) and os.path.exists(ID_MAP_PATH):
            print("📦 Loading FAISS index from disk...")
            self.index = faiss.read_index(INDEX_PATH)
            self.deck_ids = np.load(ID_MAP_PATH).tolist()
            print(f"✅ Loaded {len(self.deck_ids)} embeddings from disk.")
            return

        print("🔄 No index on disk, loading from DB...")
        records = db.fetch_all_deck_embeddings()
        if not records:
            print("⚠️ No embeddings found in DB.")
            return

        embeddings = np.array([r['embedding'] for r in records]).astype('float32')
        self.deck_ids = [r['deck_id'] for r in records]
        self.index = faiss.IndexFlatL2(embeddings.shape[1])
        self.index.add(embeddings)

        faiss.write_index(self.index, INDEX_PATH)
        np.save(ID_MAP_PATH, np.array(self.deck_ids))
        print(f"✅ Built FAISS index with {len(self.deck_ids)} entries.")

    def search(self, cards, top_k=3):
        query = " ".join(cards)
        vec = self.model.encode([query]).astype('float32')
        _, indices = self.index.search(vec, top_k)
        return [str(self.deck_ids[i]) for i in indices[0]]


class VectorDeckService(search_pb2_grpc.VectorServiceServicer):
    def __init__(self):
        self.store = FaissVectorStore()

    def SearchSimilarDecks(self, request, context):
        result = self.store.search(request.cards)
        return search_pb2.DeckResult(decks=result)

class VectorCardService(search_pb2_grpc.VectorServiceServicer):
    def __init__(self):
        self.store = FaissVectorStore()

    def SearchSimilarCards(self, request, context):
        # Not implemented for FAISS store
        return search_pb2.CardResult(cards=[])


def serve(mode):
    if mode == "card":
        vector = VectorCardService()
    elif mode == "deck":
        vector = VectorDeckService()
    else:
        raise ValueError("Invalid mode: choose 'card' or 'deck'")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))
    search_pb2_grpc.add_VectorServiceServicer_to_server(vector, server)
    server.add_insecure_port('[::]:60051')
    print("🚀 FAISS vector store gRPC server running on port 60051")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python service.py [card|deck]")
        sys.exit(1)
    serve(sys.argv[1])
