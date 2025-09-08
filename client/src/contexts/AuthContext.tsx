import { createContext, useContext, useState } from "react"
// import { useLocation } from "react-router-dom"


const Context  = createContext<{
    atk: string, 
    setAtkFunc: (newToken: string) => void
}>({
    atk: "", 
    setAtkFunc: () => {}
})
const AuthContextProvider = ({children}: {children: React.ReactNode}) => {
  const [atk, setAtk] = useState("")
  // const path = useLocation().pathname




  // useEffect(() => {
  //   const savedToken = localStorage.getItem("atk");
  //   if (savedToken) {
  //     setAtk(savedToken);
  //   }
  //   else if(!path.includes("/auth/") && !atk){
  //       // window.location.href = "/auth/login"
  //       return
  //   } 
  //   else return
  // }, [atk, path]);

  const setAtkFunc = (newToken: string) => {
    if (newToken) {
        setAtk(newToken);
        localStorage.setItem("atk", newToken);
    } else {
        localStorage.removeItem("atk");
        setAtk("");
    }
  };

  


  
  return (
    <Context.Provider value={{atk, setAtkFunc}}>
        {children}
    </Context.Provider>
  )
}

export default  AuthContextProvider

export const useAuthContext = () => {
  return useContext(Context)
}