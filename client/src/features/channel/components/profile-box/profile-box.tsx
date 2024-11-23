import { styles } from "./profile-box.css";
import { useAuthUser } from "@/hooks/auth-user";

export const ProfileBox = () => {
  const user = useAuthUser();

  if (!user.isAuthenticated || !user.data) {
    return null;
  }

  return (
    <div className={styles.profileBox}>
      <img src={user.data.imageURL} alt="icon" className={styles.iconUrl} />
      <div className={styles.profileName}>
        <span>{user.data.name}</span>
      </div>
    </div>
  );
};
