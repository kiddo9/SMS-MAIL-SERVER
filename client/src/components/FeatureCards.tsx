import { FileSpreadsheet, MessageSquare, NotepadText } from 'lucide-react'

const FeatureCards = () => {
  return (
    <div className="grid md:grid-cols-3 gap-6 mb-12">
            <div className="bg-white rounded-xl p-6 shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
              <FileSpreadsheet className="w-10 h-10 text-[#6699ff] mx-auto mb-4" />
              <h3 className="font-semibold text-gray-800 mb-2">Easy Upload</h3>
              <p className="text-sm text-gray-600">Support for Excel (.xlsx) and CSV files with student data</p>
            </div>
            
            <div className="bg-white rounded-xl p-6 shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
              <NotepadText className="w-10 h-10 text-green-600 mx-auto mb-4" />
              <h3 className="font-semibold text-gray-800 mb-2">Templating</h3>
              <p className="text-sm text-gray-600">Create personalized SMS templates</p>
            </div>
            
            <div className="bg-white rounded-xl p-6 shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
              <MessageSquare className="w-10 h-10 text-purple-600 mx-auto mb-4" />
              <h3 className="font-semibold text-gray-800 mb-2">Bulk SMS</h3>
              <p className="text-sm text-gray-600">Send reminders to multiple students instantly</p>
            </div>
          </div>
  )
}

export default FeatureCards