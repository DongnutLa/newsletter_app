export interface Newsletter {
  id: string;
  template: string;
  subject: string;
  file: string;
  recipients: string[];
  active: boolean;
  schedule?: string;
  sentCount?: number;
}

export interface CreateNewsletterDTO {
  template: string;
  file: string;
  recipients: string[];
  subject: string;
  topic: string;
}

export interface SendNewsletterDTO {
  newsletterId: string;
  extraEmail?: string;
}

export const CREATE_NEWSLETTER_DTO_INITIALS: CreateNewsletterDTO = {
  template: "",
  file: "",
  recipients: [],
  subject: "",
  topic: "",
};
