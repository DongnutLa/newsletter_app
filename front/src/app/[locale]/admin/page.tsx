import { listNewsletters } from "@/lib/services/newsletter";
import AdminContainer from "./container/AdminContainer";
import { cookies } from "next/headers";

const AdminPage = async () => {
  const newsletters = await getInitialData();
  return <AdminContainer newsletters={newsletters} />;
};

export default AdminPage;

const getInitialData = async () => {
  const token = cookies().get("_auth_")?.value;

  const list = await listNewsletters({ page: 1, pageSize: 100 }, token);
  return list.data;
};
