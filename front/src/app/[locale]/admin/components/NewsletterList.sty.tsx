import { ImageWrapper } from "@/components/global";
import { Aspect } from "@/lib/models";
import Image from "next/image";
import { CgAdd } from "react-icons/cg";
import styled from "styled-components";

export const NewsletterListContainer = styled.div`
  display: flex;
  flex-direction: column;
`;

export const NewsletterTitle = styled.h1`
  text-align: center;
`;

export const NewslettersWrapper = styled.div`
  margin-top: 28px;
  display: flex;
  flex-wrap: wrap;
  padding: 4px;
  gap: 12px;
  max-width: 1200px;
`;

export const NewsletterCard = styled.div`
  width: 200px;
  padding: 8px;
  background-color: #e9e9e9;
  border-radius: 12px;
`;

export const NewsletterImageWrapper = styled(ImageWrapper)<Aspect>`
  aspect-ratio: ${({ aspect }) => `${aspect[0]} / ${aspect[1]}`};
  margin: 0 auto;
  margin-bottom: 12px;
  width: 100%;
`;
export const NewsletterImg = styled(Image)``;

export const NewsletterBody = styled.p`
  text-align: center;
`;

export const NewsletterFooter = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding: 12px;
`;

export const NewsletterMetaWrapper = styled.div`
  display: flex;
  flex-direction: column;
`;

export const NewsletterMetadata = styled.span`
  font-size: 12px;
  color: #8a8a8a;
`;

export const NewsletterSendButton = styled.button`
  border-radius: 10px;
  color: white;
  background-color: #f53954;
  border: unset;
  width: 70px;
  height: 20px;
  cursor: pointer;

  &:hover {
    background-color: #ed6e80;
    transition: 0.3s;
  }
`;

export const AddNewsletter = styled(CgAdd)`
  color: #ed6e80;
  cursor: pointer;
  transition: 0.3s;
  &:hover {
    color: #f53954;
  }
`;
