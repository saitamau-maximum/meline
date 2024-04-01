export interface IChannelRepository {
  createChannel: (name: string) => Promise<Response>;
}

export class ChannelRepository implements IChannelRepository {
  public createChannel = async (name: string) => {
    const res = await fetch("/api/channel", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ name }),
    });

    return res;
  };
}
