
import FeatureCards from './FeatureCards'

const Hero = () => {
  return (
    <div className="text-center mb-4">
          <div className="flex justify-center items-center px-4 py-2 rounded-full text-lg font-bold mb-6 text-center">
            <img className="rotate-y-180" width={30} src="/logo-icon.png" alt="Logo"/>
            <span>SMS Notifications Made Easy</span>
            <img className="" width={30} src="/logo-icon.png" alt="Logo" />
          </div>
          
          <h2 className="text-4xl font-bold text-gray-900 mb-4">
            Automated Fee Reminder System
          </h2>
          
          <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
            Upload Neo Cloud student data file and automatically send personalized SMS reminders to students with outstanding fees. Simple, efficient, and effective.
          </p>

          {/* Feature Cards */}
          <FeatureCards/>
    </div>
  )
}

export default Hero