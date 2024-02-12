import threading
import pika

class RabbitMQConsumer:
    def __init__(self, queue_name, message_handler):
        self.queue_name = queue_name
        self.message_handler = message_handler
        self.connection = None
        self.host = None
        self.port = None
        self.username = None
        self.password = None
        self.channel = None
        self.thread = None
        self.stop_event = threading.Event()
    

    def connect(self):
        credentials = pika.PlainCredentials(self.username, self.password)
        parameters = pika.ConnectionParameters(host=self.host, port=self.port, credentials=credentials)
        self.connection = pika.BlockingConnection(parameters)
        self.channel = self.connection.channel()
        self.channel.queue_declare(queue=self.queue_name)
        print(f"Connected to RabbitMQ on {self.host}:{self.port}")

    def start_consuming(self):
        self.thread = threading.Thread(target=self._consume)
        self.thread.start()
        print(f"Started consuming messages from queue '{self.queue_name}'")

    def _consume(self):
        self.channel.basic_consume(queue=self.queue_name, on_message_callback=self._callback, auto_ack=True)
        self.channel.start_consuming()

    def _callback(self, ch, method, properties, body):
        message = body.decode()
        message_thread = threading.Thread(target=self.message_handler, args=(message,))
        message_thread.start()

    def stop_consuming(self):
        self.stop_event.set()
        if self.thread:
            self.thread.join()
        if self.connection:
            self.connection.close()
        print(f"Stopped consuming messages from queue '{self.queue_name}'")