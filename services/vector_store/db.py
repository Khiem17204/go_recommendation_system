# db.py
import psycopg2
from psycopg2.extras import RealDictCursor
import os

DB_URL = (
    f"postgresql://"
    f"{os.getenv('POSTGRES_USER', 'postgres')}:"
    f"{os.getenv('POSTGRES_PASSWORD', 'password')}@"
    f"{os.getenv('DB_HOST', 'localhost')}:"
    f"{os.getenv('DB_PORT', '5432')}/"
    f"{os.getenv('DB_NAME', 'go_rec_sys')}"
)

def get_conn():
    return psycopg2.connect(DB_URL, cursor_factory=RealDictCursor)

def fetch_unembedded_decks():
    with get_conn() as conn:
        with conn.cursor() as cur:
            cur.execute("""
                SELECT d.id, d.deck_name
                FROM decks d
                LEFT JOIN deck_embedding e ON d.id = e.deck_id
                WHERE e.deck_id IS NULL
            """)
            return cur.fetchall()

def fetch_deck_cards(deck_id):
    with get_conn() as conn:
        with conn.cursor() as cur:
            cur.execute("""
                SELECT c.name
                FROM cards_in_deck cid
                JOIN cards c ON cid.card_id = c.id
                WHERE cid.deck_id = %s
            """, (deck_id,))
            return [row['name'] for row in cur.fetchall()]

def insert_deck_embedding(deck_id, name, embedding):
    with get_conn() as conn:
        with conn.cursor() as cur:
            cur.execute("""
                INSERT INTO deck_embedding (deck_id, name, embedding)
                VALUES (%s, %s, %s)
            """, (deck_id, name, embedding.tolist()))
            conn.commit()
