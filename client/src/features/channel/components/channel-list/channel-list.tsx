import { Link } from "react-router-dom";
import { styles } from "./channel-list.css";
import { Hash } from "react-feather";
import { clsx } from "@/libs/clsx";

interface Channel {
  id: number;
  name: string;
  active?: boolean;
  hasNotification?: boolean;
}

interface Props {
  channels: Channel[];
  isLoading?: boolean;
}

export const ChannelList = ({ channels, isLoading }: Props) => {
  if (isLoading) {
    return (
      <div className={styles.channelList}>
        {Array.from({ length: 5 }).map((_, index) => (
          <div key={index} className={styles.channelListItemSkeleton} />
        ))}
      </div>
    );
  }

  return (
    <div className={styles.channelList}>
      {channels.map((channel) => (
        <Link
          key={channel.id}
          className={clsx(
            styles.channelListItem,
            channel.hasNotification && styles.channelListItemNotification,
            channel.active && styles.channelListItemActive
          )}
          to={`/channel/${channel.id}`}
        >
          <Hash className={styles.channelListItemIcon} />
          {channel.name}
        </Link>
      ))}
    </div>
  );
};
