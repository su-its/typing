import { Flex, VStack } from "@chakra-ui/react";
import GameStartButton from "../molecules/GameStartButton";
import RankingButton from "../molecules/RankingButton";
import LogoutButton from "../molecules/LogoutButton";

const HomeMenuContainer = () => {
  return (
    <Flex justify="center" align="center" h="80vh">
      <VStack>
        {/* TODO: ログイン状況に応じて表示を切り替え */}
        <GameStartButton />
        <RankingButton />
        <LogoutButton />
      </VStack>
    </Flex>
  );
};

export default HomeMenuContainer;
