import { GithubIcon } from "@/components/icons/github";
import { styles } from "./login.css";

export const Login = () => {
  return (
    <div className={styles.container}>
    <a className={styles.content} href="/api/auth/login">
      <GithubIcon />
      <span className={styles.loginLabel}>Githubでログイン</span>
    </a>
    </div>
  );
};
