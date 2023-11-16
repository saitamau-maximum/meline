import { GithubIcon } from "@/components/icons/github";
import { styles } from "./auth-pane.css";

export const AuthPane = () => {
  return (
    <div className={styles.container}>
      <a className={styles.content} href="/api/auth/login">
        <GithubIcon />
        <span className={styles.loginLabel}>Githubでログイン</span>
      </a>
    </div>
  );
};
