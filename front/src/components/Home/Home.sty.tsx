import styled from "styled-components";

export const HomeWrapper = styled.main`
  display: grid;
  place-content: center;
  height: 100vh;
  width: 100vw;
  background-color: #f53954;
`;

export const NewsletterWrapper = styled.div`
  width: 500px;
  display: flex;
  flex-direction: column;
  gap: 20px;
`;

export const NewsletterTitle = styled.h1`
  color: white;
`;

export const NewsletterInput = styled.div`
  background-color: white;
  display: flex;
  justify-content: space-between;
  padding: 4px;
  height: 40px;
  border-radius: 40px;
`;

export const NewsletterEmail = styled.input`
  border: unset;
  border-radius: 40px;
  &:focus-visible {
    outline: unset;
  }
  margin-left: 12px;
  width: 90%;
`;

export const NewsletterSubmit = styled.button`
  border-radius: 40px;
  color: white;
  background-color: #f53954;
  border: unset;
  width: 100px;
  cursor: pointer;

  &:hover {
    background-color: #ed6e80;
    transition: 0.3s;
  }

  &:disabled {
    background-color: #ddc1c5;
    cursor: not-allowed;
  }
`;
