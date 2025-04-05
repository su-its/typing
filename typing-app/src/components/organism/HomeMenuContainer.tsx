import GameStartButton from "../molecules/GameStartButton";
import RankingButton from "../molecules/RankingButton";
import LogoutButton from "../molecules/LogoutButton";
import styles from "@/assets/sass/organism/HomeMenuContainer.module.scss";

const HomeMenuContainer = () => {
  return (
    <div className={styles.menu}>
      <div className={styles.container}>
        {/* TODO: ログイン状況に応じて表示を切り替え */}
        <GameStartButton />
        <RankingButton />
        <LogoutButton />
      </div>
    </div>
  );
};

export default HomeMenuContainer;
