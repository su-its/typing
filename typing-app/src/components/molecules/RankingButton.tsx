"use client";

import { useRouter } from "next/navigation";
import { Button } from "@chakra-ui/react";

const RankingButton = () => {
  const router = useRouter();

  const handleRouteRanking = async () => {
    //TODO:ログインの維持を実装
    router.push("/ranking");
  };

  return (
    <Button colorScheme="orange" size="lg" onClick={handleRouteRanking}>
      Ranking
    </Button>
  );
};

export default RankingButton;
