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
}

export const ChannelList = ({ channels }: Props) => {
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
