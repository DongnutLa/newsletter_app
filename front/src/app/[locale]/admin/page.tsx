import { cookies } from "next/headers";
import AdminContainer from "./container/AdminContainer";
import { redirect } from "next/navigation";

const AdminPage = async () => {
  const token = cookies().get("_auth_")?.value;
  if (!token) {
    redirect("/login");
  }

  return <AdminContainer />;
};

export default AdminPage;
