

const NavBar = () => {
  return (
    <nav className='bg-white shadow-2xs sticky top-0 z-30'>
        <div className="flex justify-between items-center py-1 px-10">
          <div className="flex items-center">
            <img width={150} src="/logo.png" alt="Company Logo" className="mr-2" />
            <h1 className="text-2xl font-bold text-[#6699ff]">SMS Service</h1>
          </div>
          <div className="flex items-center">
            <a href="/templates" className="text-[#6699ff] hover:text-blue-800 mr-4">Templates</a>
            <a href="/sent" className="text-[#6699ff] hover:text-blue-800 mr-4">Sent</a>
            <a href="#contact" className="text-[#6699ff] hover:text-blue-800">Contact</a>
          </div>
        </div>
      </nav>
  )
}

export default NavBar