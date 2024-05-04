import { Message } from "./message";

export interface MessageResponse {
  message: Message;
}

export interface IChatRepository {
  connect: () => void;
  disconnect: () => void;
  onMessageReceived: (callback: (res: MessageResponse) => void) => void;
}

export class ChatRepositoryImpl implements IChatRepository {
  private connection: WebSocket | null = null;
  private channelId: number;

  constructor(channelId: number) {
    this.channelId = channelId;
  }

  connect() {
    const protocol = location.protocol === "https:" ? "wss" : "ws";
    const host = location.host;
    const url = `${protocol}://${host}/api/ws/${this.channelId}`;
    this.connection = new WebSocket(url);
    this.connection?.addEventListener("close", () => {
      this.connect();
    });
  }

  disconnect() {
    this.connection?.close();
  }

  onMessageReceived(callback: (res: MessageResponse) => void) {
    this.connection?.addEventListener("message", (event) => {
      const res = JSON.parse(event.data);
      callback(res);
    });
  }
}
