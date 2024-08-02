def create_id_mapping(unique_ids):
    return {id_value: index for index, id_value in enumerate(unique_ids)}

def map_ids_to_indices(ids, id_to_index_mapping):
    return [id_to_index_mapping.get(id_value, -1) for id_value in ids]
