import { MaximumIcon } from "@/components/icons/maximum";
import { styles } from "./layout.css";
import { CreateChannelForm } from "../create-channel-form";
import { ChannelList } from "../channel-list";
import { Outlet } from "react-router-dom";

interface ChannelLayoutProps {
  channels: {
    id: number;
    name: string;
  }[];
  isChannelsLoading?: boolean;
}

export const ChannelLayout = ({
  channels,
  isChannelsLoading,
}: ChannelLayoutProps) => {
  return (
    <div className={styles.channelLayoutWrapper}>
      <div className={styles.channelLayoutSidePanel}>
        <div className={styles.channelLayoutSidePanelLogo}>
          <MaximumIcon />
        </div>
        <CreateChannelForm />
        <ChannelList channels={channels} isLoading={isChannelsLoading} />
      </div>
      <div className={styles.channelLayoutMain}>
        <Outlet />
      </div>
    </div>
  );
};
