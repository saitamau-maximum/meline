import { Message } from "./message";

export interface MessageResponse {
  message: Message;
}

export interface IChatRepository {
  connect: (channelId: number) => void;
  disconnect: () => void;
  onMessageReceived: (callback: (res: MessageResponse) => void) => void;
}

export class ChatRepositoryImpl implements IChatRepository {
  private connection: WebSocket | null = null;

  connect(channelId: number) {
    const protocol = location.protocol === "https:" ? "wss" : "ws";
    const host = location.host;
    const url = `${protocol}://${host}/api/ws/${channelId}`;
    this.connection = new WebSocket(url);
    this.connection?.addEventListener("close", () => {
      this.connect(channelId);
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
