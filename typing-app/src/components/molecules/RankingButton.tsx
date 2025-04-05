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
    <div className={styles.button} onClick={handleRouteRanking}>
      <img src={rankingButton.src} />
    </div>
  );
};

export default RankingButton;
