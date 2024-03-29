import { Flex, VStack } from "@chakra-ui/react";
import GameStartButton from "../molecules/GameStartButton";
import RankingButton from "../molecules/RankingButton";
import LogoutButton from "../molecules/LogoutButton";

const HomeMenuContainer = () => {
  return (
    <Flex justify="center" align="center" h="65vh">
      <VStack spacing={8} align="stretch" width="50%" maxWidth="md" mx="auto">
        {/* TODO: ログイン状況に応じて表示を切り替え */}
        <GameStartButton />
        <RankingButton />
        <LogoutButton />
      </VStack>
    </Flex>
  );
};

export default HomeMenuContainer;
