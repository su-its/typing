"use client";

import { useRouter } from "next/navigation";
import rankingButton from "@/assets/images/home/ranking.png";
import styles from "@/assets/sass/molecules/RankingButton.module.scss";

const RankingButton = () => {
  const router = useRouter();

  const handleRouteRanking = async () => {
    //TODO:ログインの維持を実装
    router.push("/ranking");
  };

  return (
    <div className={styles.button} onClick={handleRouteRanking} aria-label="see ranking">
      {/* eslint-disable-next-line @next/next/no-img-element */}
      <img src={rankingButton.src} alt="" />
    </div>
  );
};

export default RankingButton;
