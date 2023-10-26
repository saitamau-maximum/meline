import { clsx } from "@/libs/clsx";
import { styles } from "./loading-overlay.css";
import { LoadingOverlayContext } from "@/providers/loading-overlay";
import { vars } from "@/styles/theme.css";
import { useContext, useEffect, useRef } from "react";

export const LoadingOverlay = () => {
  const { isLoading } = useContext(LoadingOverlayContext);
  const overlayRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!isLoading) {
      overlayRef.current?.classList.add(styles.fade);
    } else {
      overlayRef.current?.classList.remove(styles.fade);
    }
  }, [isLoading]);

  return (
    <div
      ref={overlayRef}
      className={clsx(styles.overlay, {
        [styles.active]: isLoading,
      })}
    >
      <img src="/maximum.svg" alt="Maximum" width="300" height="100" />
      <svg
        width="38"
        height="38"
        viewBox="0 0 38 38"
        xmlns="http://www.w3.org/2000/svg"
        className={styles.loading}
      >
        <defs>
          <linearGradient x1="8.042%" y1="0%" x2="65.682%" y2="23.865%" id="a">
            <stop
              stop-color={vars.color.gray[8]}
              stop-opacity="0"
              offset="0%"
            />
            <stop
              stop-color={vars.color.gray[8]}
              stop-opacity=".631"
              offset="63.146%"
            />
            <stop stop-color={vars.color.gray[8]} offset="100%" />
          </linearGradient>
        </defs>
        <g fill="none" fill-rule="evenodd">
          <g transform="translate(1 1)">
            <path
              d="M36 18c0-9.94-8.06-18-18-18"
              id="Oval-2"
              stroke="url(#a)"
              stroke-width="2"
            >
              <animateTransform
                attributeName="transform"
                type="rotate"
                from="0 18 18"
                to="360 18 18"
                dur="0.9s"
                repeatCount="indefinite"
              />
            </path>
            <circle fill={vars.color.gray[8]} cx="36" cy="18" r="1">
              <animateTransform
                attributeName="transform"
                type="rotate"
                from="0 18 18"
                to="360 18 18"
                dur="0.9s"
                repeatCount="indefinite"
              />
            </circle>
          </g>
        </g>
      </svg>
    </div>
  );
};
