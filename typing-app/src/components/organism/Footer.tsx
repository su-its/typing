import React from "react";
import BrandText from "../molecules/BrandText";
import Separator from "../atoms/Separater";
import styles from "@/assets/sass/organism/Footer.module.scss";
import soundOnImage from "@/assets/images/soundon.svg";
import soundOffImage from "@/assets/images/soundoff.svg";

const date = new Date();

type Props = {
  isPlay: boolean;
  setIsPlay: React.Dispatch<React.SetStateAction<boolean>>;
};

const Footer: React.FC<Props> = ({ isPlay, setIsPlay }) => {
  const toggleSound = () => {
    setIsPlay((prev) => !prev);
  };
  return (
    <div className={styles.footer}>
      <Separator />
      <div className={styles.left}>
        <BrandText />
        <div className={styles.text}>
          (c) 2024-{date.getFullYear()} Faculty of Informatics, Shizuoka University all rights reserved. Developed by IT
          Solution Room, Shizuoka University.
        </div>
      </div>
      <div className={styles.right}>
        <div className={styles.sound} onClick={toggleSound}>
          <img src={isPlay ? soundOnImage.src : soundOffImage.src} alt={isPlay ? "SOUND ON" : "SOUND OFF"} />
        </div>
      </div>
    </div>
  );
};

export default Footer;
