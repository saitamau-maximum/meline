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

export const MessageRepositoryImpl: IMessageRepository = {
  createMessage: async (channelId, param) => {
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
  },
  getMessages: async (channelId) => {
    const res = await fetch(`/api/channel/${channelId}/message`);

    return res.json();
  },
  getMessages$$key: (channelId) => ["getMessages", `channelId:${channelId}`],
};
