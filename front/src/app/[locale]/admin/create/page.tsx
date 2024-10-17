import { listUsers } from "@/lib/services/users";
import CreateNewsletterContainer from "../container/CreateNewsletterContainer";
import { cookies } from "next/headers";

const CreatePage = async () => {
  const users = await getInitialData();
  return <CreateNewsletterContainer users={users} />;
};

export default CreatePage;

const getInitialData = async () => {
  const token = cookies().get("_auth_")?.value;

  return await listUsers(token);
};
