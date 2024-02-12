import os

def env_mongo_uri():
    return os.getenv("MONGO_URI")

def env_db_name():
    return os.getenv("DB_NAME")

def env_request_collection():
    return os.getenv("REQUEST_COLLECTION")

def env_user_collection():
    return os.getenv("USER_COLLECTION")

def env_training_run_collection():
    return os.getenv("TRAINING_RUNS_COLLECTION")

def env_exporter_collection():
    return os.getenv("EXPORTER_COLLECTION")

def env_exporter_runs_collection():
    return os.getenv("EXPORTER_RUNS_COLLECTION")

def env_dataset_collection():
    return os.getenv("DATASET_COLLECTION")

def env_admin_email():
    return os.getenv("EMAIL")

def env_minio_uri():
    return os.getenv("MINIO_URI")

def env_minio_access_key():
    return os.getenv("MINIO_ACCESS_KEY")

def env_minio_private_key():
    return os.getenv("MINIO_PRIVATE_KEY")

def env_export_bucket_name():
    return os.getenv("EXPORT_BUCKET_NAME")

def env_model_bucket_name():
    return os.getenv("MODEL_BUCKET_NAME")

def env_extractor_bucket_name():
    return os.getenv("EXTRACTOR_BUCKET_NAME")

def env_admin_api_key():
    return os.getenv("ADMIN_API_KEY")

def env_model_collection():
    return os.getenv("MODELS_COLLECTION")

def env_mq_uri():
    return os.getenv("RABBIT_URI")

def env_train_queue_name():
    return os.getenv("TRAIN_QUEUE")

def env_export_queue_name():
    return os.getenv("EXPORT_QUEUE")
