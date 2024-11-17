export interface CreateMessageParam {
  content: string;
}

interface User {
  id: number;
  name: string;
  image_url: string;
}

interface ReplyToMessage {
  id: string;
  user: User;
  content: string;
}

export interface Message {
  id: string;
  user: User;
  content: string;
  reply_to_message: ReplyToMessage | null;
  created_at: string;
  updated_at: string;
}

interface CreateMessageResponse {
  messages: Message[];
}

export interface IMessageRepository {
  createMessage: (
    channelId: number,
    param: CreateMessageParam
  ) => Promise<void>;
  getMessages: (channelId: number) => Promise<CreateMessageResponse>;
  getMessages$$key: (channelId: number) => string[];
}

export class MessageRepositoryImpl implements IMessageRepository {
  async createMessage(channelId: number, param: CreateMessageParam) {
    const res = await fetch(`/api/channel/${channelId}/message`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(param),
    });

    if (!res.ok) {
      throw new Error("Failed to create message");
    }
  }

  async getMessages(channelId: number) {
    const res = await fetch(`/api/channel/${channelId}/message`);

    return res.json();
  }

  getMessages$$key(channelId: number) {
    return ["getMessages", `channelId:${channelId}`];
  }
}
