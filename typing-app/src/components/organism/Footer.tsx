import React from "react";
import BrandText from "../molecules/BrandText";
import Separator from "../atoms/Separater";
import styles from "@/assets/sass/organism/Footer.module.scss";
import soundOnImage from "@/assets/images/soundon.svg";
import soundOffImage from "@/assets/images/soundoff.svg";

const date = new Date();

type Props = {
  isMuted: boolean;
  setIsMuted: React.Dispatch<React.SetStateAction<boolean>>;
};

const Footer: React.FC<Props> = ({ isMuted, setIsMuted }) => {
  const toggleSound = () => {
    setIsMuted((prev) => !prev);
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
          <img src={isMuted ? soundOffImage.src : soundOnImage.src} alt={isMuted ? "SOUND OFF" : "SOUND ON"} />
        </div>
      </div>
    </div>
  );
};

export default Footer;
