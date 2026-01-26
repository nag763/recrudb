import { useUser } from "@/hooks/useAuth";
import { LoginComponent } from "./LoginComponent";

export default function AuthGuard({ children }: { children: React.ReactNode }) {
    const { data: user, isLoading } = useUser();

    if(isLoading) {
        return <p>Waiting...</p>
    }

    if(!user) {
        return <LoginComponent/>
    }

    return <>{children}</>
}