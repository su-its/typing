"use client";

import { useRouter } from "next/navigation";
import { Image } from "@chakra-ui/react";
import rankingButton from "@/assets/images/home/ranking.png";

const RankingButton = () => {
  const router = useRouter();

  const handleRouteRanking = async () => {
    //TODO:ログインの維持を実装
    router.push("/ranking");
  };

  return <Image mb={2} src={rankingButton.src} onClick={handleRouteRanking} cursor="pointer" />;
};

export default RankingButton;
