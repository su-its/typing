import { CustomButton } from "../atoms/CustomButton";

interface PaginationProps {
  onPrev: () => void;
  onNext: () => void;
  isPrevDisabled: boolean;
  isNextDisabled: boolean;
}

export const Pagination = ({ onPrev, onNext, isPrevDisabled, isNextDisabled }: PaginationProps) => (
  <div>
    <CustomButton onClick={onPrev} isDisabled={isPrevDisabled}>
      前のページ
    </CustomButton>
    <CustomButton onClick={onNext} isDisabled={isNextDisabled}>
      次のページ
    </CustomButton>
  </div>
);
