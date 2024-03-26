import { Flex } from "@chakra-ui/react";
import { CustomButton } from "../atoms/CustomButton";

interface PaginationProps {
  onPrev: () => void;
  onNext: () => void;
  isPrevDisabled: boolean;
  isNextDisabled: boolean;
}

export const Pagination = ({ onPrev, onNext, isPrevDisabled, isNextDisabled }: PaginationProps) => (
  <Flex>
    <CustomButton onClick={onPrev} isDisabled={isPrevDisabled}>
      Prev
    </CustomButton>
    <CustomButton onClick={onNext} isDisabled={isNextDisabled}>
      Next
    </CustomButton>
  </Flex>
);
