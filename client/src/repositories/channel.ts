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

export class ChannelRepositoryImpl implements IChannelRepository {
  async createChannel(param: CreateChannelParam) {
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
  }

  async getJoinedChannels() {
    const res = await fetch("/api/channel");

    return res.json();
  }

  getJoinedChannels$$key() {
    return ["getJoinedChannels"];
  }
}
