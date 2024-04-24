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
  isChannelsLoading?: boolean;
}

export const ChannelTopPageTemplate = ({
  user,
  channels,
  isChannelsLoading,
}: ChannelTopPageProps) => {
  return (
    <ChannelLayout
      sidePanel={
        <>
          <CreateChannelForm />
          <ChannelList channels={channels} isLoading={isChannelsLoading} />
        </>
      }
      main={<div>{JSON.stringify(user)}</div>}
    />
  );
};
