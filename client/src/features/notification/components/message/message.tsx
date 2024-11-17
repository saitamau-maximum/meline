import { styles } from "./message.css";

interface MessageProps {
  avatarUrl: string;
  username: string;
  message: string;
}

export const Message = ({ avatarUrl, username, message }: MessageProps) => {
  return (
    <div className={styles.container}>
      <img
        src={avatarUrl}
        alt={`${username}のアイコン`}
        className={styles.avatar}
        width="40"
        height="40"
      />
      <div className={styles.metaContainer}>
        <span className={styles.username}>{username}</span>
        <p className={styles.message}>{message}</p>
      </div>
    </div>
  );
};
