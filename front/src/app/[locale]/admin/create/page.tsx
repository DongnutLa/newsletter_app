import { cookies } from "next/headers";
import CreateNewsletterContainer from "../container/CreateNewsletterContainer";
import { redirect } from "next/navigation";

const CreatePage = async () => {
  const token = cookies().get("_auth_")?.value;
  if (!token) {
    redirect("/login");
  }

  return <CreateNewsletterContainer />;
};

export default CreatePage;
