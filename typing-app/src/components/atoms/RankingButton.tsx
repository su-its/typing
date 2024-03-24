// components/buttons/RankingButton.js
"use client";

import Link from "next/link";
import { Button } from "@chakra-ui/react";

const RankingButton = () => {
  return (
    <Button as={Link} href="/ranking" colorScheme="orange" size="lg">
      Ranking
    </Button>
  );
};

export default RankingButton;
