from sqlalchemy import create_engine, text
import pandas as pd
import torch
from torch.utils.data import TensorDataset, ConcatDataset
from torch.utils.data import DataLoader
import numpy as np
from scipy.sparse import dok_matrix

# Database connection setup
engine = create_engine('postgresql://postgres:cghuy2004@localhost:5432/go_rec_sys')

def get_data_iterator(chunk_size=10000):
    """Fetches data from the database in chunks."""
    offset = 0
    with engine.connect() as connection:
        while True:
            query = text("""
                SELECT card_id, deck_id, card_count
                FROM cards_in_deck
                ORDER BY card_id
                LIMIT :chunk_size OFFSET :offset
            """)
            result = connection.execute(query, {"chunk_size": chunk_size, "offset": offset})
            chunk = pd.DataFrame(result.fetchall(), columns=result.keys())
            if chunk.empty:
                break
            yield chunk
            offset += chunk_size

def process_chunk(chunk):
    """Converts a dataframe chunk into a PyTorch TensorDataset."""
    cards = torch.LongTensor(chunk['card_id'].values)
    decks = torch.LongTensor(chunk['deck_id'].values)
    num_card = torch.FloatTensor(chunk['card_count'].values)
    return TensorDataset(decks, cards, num_card)

# Function to get unique card and deck IDs
def get_unique_ids():
    card_ids = set()
    deck_ids = set()
    for chunk in get_data_iterator():
        card_ids.update(chunk['card_id'].unique())
        deck_ids.update(chunk['deck_id'].unique())
    return list(card_ids), list(deck_ids)

# Retrieve all card IDs from the 'cards' table
def get_all_card_ids():
    with engine.connect() as connection:
        query = text("SELECT DISTINCT id FROM cards")
        result = connection.execute(query)
        all_card_ids = [row[0] for row in result.fetchall()]  # Use index 0 to access the first element in each tuple
    return all_card_ids


# Identify cards not present in any decks
def get_missing_card_ids(existing_card_ids):
    all_card_ids = set(get_all_card_ids())
    existing_card_ids = set(existing_card_ids)
    missing_card_ids = all_card_ids - existing_card_ids
    return list(missing_card_ids)

# Add missing cards to the dataset
def add_missing_cards_to_dataset(missing_card_ids, datasets):
    """Add missing cards with zero counts to the dataset."""
    placeholder_deck_id = -1  # Placeholder deck ID for missing cards
    placeholder_count = 0     # Count is zero since the card is not in any deck
    
    for card_id in missing_card_ids:
        card_tensor = torch.tensor([card_id], dtype=torch.long)
        deck_tensor = torch.tensor([placeholder_deck_id], dtype=torch.long)
        count_tensor = torch.tensor([placeholder_count], dtype=torch.float)
        datasets.append(TensorDataset(deck_tensor, card_tensor, count_tensor))

# Process all chunks and add missing cards
datasets = []
for chunk in get_data_iterator():
    dataset = process_chunk(chunk)
    datasets.append(dataset)

# Get unique card and deck IDs
unique_card_ids, unique_deck_ids = get_unique_ids()

# Get the list of missing card IDs
missing_card_ids = get_missing_card_ids(unique_card_ids)

# Add missing cards to the dataset list
add_missing_cards_to_dataset(missing_card_ids, datasets)

# Combine all datasets into a full dataset
full_dataset = ConcatDataset(datasets)

# Create DataLoader
dataloader = DataLoader(full_dataset, batch_size=32, shuffle=True)

# Utility matrix dimensions
num_decks = len(unique_deck_ids)
num_cards = len(unique_card_ids)


for i, (decks, cards, counts) in enumerate(dataloader):
    print("Decks:", decks)
    print("Cards:", cards)
    print("Counts:", counts)

