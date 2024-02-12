from dataclasses import dataclass
from typing import List, Dict, Any

@dataclass
class Model:
    ID: str
    Name: str
    Description: str
    Dims: List[int]
    TensorflowLayers: Dict[str, Any]

    def to_json(self):
        return {
            '_id': self.ID,
            'name': self.Name,
            'description': self.Description,
            'dims': self.Dims,
            'tfLayers': self.TensorflowLayers
        }

    @staticmethod
    def from_json(json):
        return Model(
            ID=json['_id'],
            Name=json['name'],
            Description=json['description'],
            Dims=json['dims'],
            TensorflowLayers=json['tfLayers']
        )