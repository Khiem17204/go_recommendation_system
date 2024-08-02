import torch
import torch.nn as nn
from data.load_data import create_id_mappings, load_data

# Set device
device = torch.device("cuda" if torch.cuda.is_available() else "cpu")

# Get mappings and counts
deck_id_map, card_id_map, num_decks, num_cards = create_id_mappings()
dataloader = load_data(deck_id_map, card_id_map)
for user_batch, item_batch, count_batch in dataloader:
    print(user_batch,item_batch,count_batch)
# Model definition as before
class NCF(nn.Module):
    def __init__(self, num_users, num_items, embedding_size, layers):
        super(NCF, self).__init__()
        self.user_embedding = nn.Embedding(num_users, embedding_size)
        self.item_embedding = nn.Embedding(num_items, embedding_size)
        self.fc_layers = nn.ModuleList()
        input_size = embedding_size * 2
        for output_size in layers:
            self.fc_layers.append(nn.Linear(input_size, output_size))
            self.fc_layers.append(nn.ReLU())
            input_size = output_size
        self.output_layer = nn.Linear(layers[-1], 4)
        self.softmax = nn.Softmax(dim=1)

    def forward(self, user_indices, item_indices):
        user_embedded = self.user_embedding(user_indices)
        item_embedded = self.item_embedding(item_indices)
        x = torch.cat([user_embedded, item_embedded], dim=-1)
        for layer in self.fc_layers:
            x = layer(x)
        logits = self.output_layer(x)
        predictions = self.softmax(logits)
        return predictions

# Initialize the model

embedding_size = 64
layers = [128, 64, 32]
num_epochs = 10
model = NCF(num_decks, num_cards, embedding_size, layers).to(device)
criterion = nn.CrossEntropyLoss()
optimizer = torch.optim.Adam(model.parameters(), lr=0.001)

# Training loop

for epoch in range(num_epochs):
    total_loss = 0
    for user_batch, item_batch, count_batch in dataloader:
        user_batch = user_batch.to(device)
        item_batch = item_batch.to(device)
        count_batch = count_batch.to(device)
        
        predictions = model(user_batch, item_batch)
        loss = criterion(predictions, count_batch)
        
        optimizer.zero_grad()
        loss.backward()
        optimizer.step()
        
        total_loss += loss.item()
    
    avg_loss = total_loss / len(dataloader)
    print(f"Epoch {epoch+1}/{num_epochs}, Loss: {avg_loss:.4f}")