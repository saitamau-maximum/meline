import { ChannelLayout } from "../../components/layout";
import { ChannelList } from "./components/channel-list";

interface ChannelTopPageProps {
  user: {
    name: string;
    imageURL: string;
  };
  channels: {
    id: number;
    name: string;
  }[];
}

export const ChannelTopPageTemplate = ({
  user,
  channels,
}: ChannelTopPageProps) => {
  return (
    <ChannelLayout
      sidePanel={
        <>
          <ChannelList channels={channels} />
        </>
      }
      main={<div>{JSON.stringify(user)}</div>}
    />
  );
};
