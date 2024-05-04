interface Channel {
  id: number;
  name: string;
}

interface GetJoinedChannelsResponse {
  channels: Channel[];
}

interface CreateChannelParam {
  name: string;
}

export interface IChannelRepository {
  createChannel: (param: CreateChannelParam) => Promise<void>;
  getJoinedChannels: () => Promise<GetJoinedChannelsResponse>;
  getJoinedChannels$$key: () => string[];
}

export const ChannelRepositoryImpl: IChannelRepository = {
  createChannel: async (param) => {
    const res = await fetch("/api/channel", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(param),
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
