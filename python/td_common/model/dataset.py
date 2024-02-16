from dataclasses import dataclass
from typing import Optional, List

@dataclass
class ReducerMetric:
    Reducer: str
    Total: int
    Tracker: int
    NonTracker: int

    def to_json(self):
        return {
            'reducer': self.Reducer,
            'total': self.Total,
            'tracker': self.Tracker,
            'nonTracker': self.NonTracker
        }
    
    @staticmethod
    def from_json(json):
        return ReducerMetric(
            Reducer=json['reducer'],
            Total=json['total'],
            Tracker=json['tracker'],
            NonTracker=json['nonTracker']
        )

@dataclass
class DataSetMetrics:
    Total: int
    ReducerMetric: List[ReducerMetric]

    def to_json(self):
        return {
            'total': self.Total,
            'reducerMetric': [rm.__dict__ for rm in self.ReducerMetric]
        }
    
    @staticmethod
    def from_json(json):
        if json is None:
            return None
        return DataSetMetrics(
            Total=json['total'],
            ReducerMetric=[ReducerMetric.from_json(rm) for rm in json['reducerMetric']]
        )

@dataclass
class Dataset:
    Name: str  # Required argument
    Description: str  # Required argument
    ID: Optional[str] = None
    Label: Optional[str] = None  # Optional argument with default value
    Metrics: Optional[DataSetMetrics] = None

    def to_json(self):
        return {
            '_id': self.ID,
            'name': self.Name,
            'label': self.Label,
            'description': self.Description,
            'metrics': self.Metrics.to_json() if self.Metrics else None
        }
    
    @staticmethod
    def from_json(json):
        return Dataset(
            ID=json['_id'],
            Name=json['name'],
            Label=json['label'],
            Description=json['description'],
            Metrics=DataSetMetrics.from_json(json['metrics']) if 'metrics' in json else None
        )