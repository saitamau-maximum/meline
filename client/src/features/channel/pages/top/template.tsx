import { IChannelRepository } from "@/repositories/channel";
import { ChannelLayout } from "../../components/layout";
import { ChannelList } from "./components/channel-list";
import { CreateChannelForm } from "./components/create-channel-form";

interface ChannelTopPageProps {
  user: {
    name: string;
    imageURL: string;
  };
  channels: {
    id: number;
    name: string;
  }[];
  fetchJoinedChannels: () => Promise<void>;
  channelRepository: IChannelRepository;
}

export const ChannelTopPageTemplate = ({
  user,
  channels,
  fetchJoinedChannels,
  channelRepository,
}: ChannelTopPageProps) => {
  return (
    <ChannelLayout
      sidePanel={
        <>
          <CreateChannelForm
            fetchJoinedChannels={fetchJoinedChannels}
            channelRepository={channelRepository}
          />
          <ChannelList channels={channels} />
        </>
      }
      main={<div>{JSON.stringify(user)}</div>}
    />
  );
};
