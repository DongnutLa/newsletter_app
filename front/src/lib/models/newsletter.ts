export interface Newsletter {
  id: string;
  template: string;
  subject: string;
  file: string;
  recipients: string[];
  active: boolean;
  schedule?: string;
}

export interface CreateNewsletterDTO {
  template: string;
  file: string;
  recipients: string[];
  subject: string;
}

export interface SendNewsletterDTO {
  newsletterId: string;
  extraEmail?: string;
}
