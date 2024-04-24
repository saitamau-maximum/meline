interface Channel {
  id: number;
  name: string;
}

interface GetJoinedChannelsResponse {
  channels: Channel[];
}

export interface IChannelRepository {
  createChannel: (name: string) => Promise<void>;
  getJoinedChannels: () => Promise<GetJoinedChannelsResponse>;
  getJoinedChannels$$key: () => string[];
}

export const ChannelRepositoryImpl: IChannelRepository = {
  createChannel: async (name: string) => {
    const res = await fetch("/api/channel", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name }),
    });

    if (!res.ok) {
      throw new Error("Failed to create channel");
    }
  },
  getJoinedChannels: async () => {
    const res = await fetch("/api/channel");

    return res.json();
  },
  getJoinedChannels$$key: () => ["getJoinedChannels"],
};
