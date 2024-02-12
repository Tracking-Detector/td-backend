from dataclasses import dataclass, field
from typing import Optional, Dict

@dataclass
class TrainingRun:
    ID: str
    ModelId: str
    Name: str
    DataSet: str
    Time: str
    F1Train: float
    F1Test: float
    TrainingHistory: Dict[str, str] = field(default_factory=dict)
    BatchSize: int
    Epochs: int

    def to_json(self):
        return {
            '_id': self.ID,
            'modelId': self.ModelId,
            'name': self.Name,
            'dataSet': self.DataSet,
            'time': self.Time,
            'f1Train': self.F1Train,
            'f1Test': self.F1Test,
            'trainingHistory': self.TrainingHistory,
            'batchSize': self.BatchSize,
            'epochs': self.Epochs
        }
    
    @staticmethod
    def from_json(json):
        return TrainingRun(
            ID=json['_id'],
            ModelId=json['modelId'],
            Name=json['name'],
            DataSet=json['dataSet'],
            Time=json['time'],
            F1Train=json['f1Train'],
            F1Test=json['f1Test'],
            TrainingHistory=json['trainingHistory'],
            BatchSize=json['batchSize'],
            Epochs=json['epochs']
        )