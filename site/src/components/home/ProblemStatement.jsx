import React from 'react';
import { Card } from "@/components/ui/card";
import { FolderSearch, Globe, ArrowRight } from 'lucide-react';

export default function ProblemStatement() {
  return (
    <div className="max-w-7xl mx-auto px-6 py-16">
      <div className="max-w-4xl mx-auto">
        {/* Problem Statement Card */}
        <Card className="bg-gray-900/50 border-gray-700 p-8 mb-8 backdrop-blur-xl">
          <div className="text-center">
            <h2 className="text-2xl md:text-3xl font-bold text-white mb-4">
              The daily developer struggle
            </h2>
            <p className="text-lg md:text-xl text-gray-300 leading-relaxed">
              You want a tiny CLI that removes the{' '}
              <span className="font-semibold text-orange-400">"where should I clone this?"</span> decision 
              and the browser tab hunt. gitr computes the path from the repo URL, clones there, 
              and opens the exact page you want with one command.
            </p>
          </div>
        </Card>

        {/* Visual Problem/Solution Flow */}
        <div className="grid md:grid-cols-3 gap-6 items-center">
          {/* Problem */}
          <Card className="bg-red-950/20 border-red-800/50 p-6 text-center">
            <div className="w-12 h-12 bg-red-600/20 rounded-xl flex items-center justify-center mx-auto mb-4">
              <FolderSearch className="w-6 h-6 text-red-400" />
            </div>
            <h3 className="text-lg font-semibold text-white mb-2">The Problem</h3>
            <p className="text-red-200 text-sm">
              "Where should I clone this repo? What was that GitHub URL again?"
            </p>
          </Card>

          {/* Arrow */}
          <div className="flex justify-center">
            <ArrowRight className="w-8 h-8 text-gray-500 hidden md:block" />
            <div className="w-8 h-0.5 bg-gray-500 md:hidden" />
          </div>

          {/* Solution */}
          <Card className="bg-green-950/20 border-green-800/50 p-6 text-center">
            <div className="w-12 h-12 bg-green-600/20 rounded-xl flex items-center justify-center mx-auto mb-4">
              <Globe className="w-6 h-6 text-green-400" />
            </div>
            <h3 className="text-lg font-semibold text-white mb-2">The Solution</h3>
            <p className="text-green-200 text-sm">
              One command: deterministic paths + instant web navigation
            </p>
          </Card>
        </div>

        {/* Quick Examples */}
        <div className="mt-12 bg-gray-900/50 border border-gray-700 rounded-2xl p-6">
          <h3 className="text-xl font-semibold text-white mb-4 text-center">
            Stop overthinking. Start coding.
          </h3>
          <div className="grid md:grid-cols-2 gap-6">
            <div>
              <h4 className="text-red-400 font-medium mb-2">❌ Before gitr:</h4>
              <div className="space-y-2 text-sm text-gray-400">
                <div>• "Hmm, where should I put this repo?"</div>
                <div>• "Let me create a new folder..."</div>
                <div>• "Now let me find the GitHub page..."</div>
                <div>• "Where did I put that repo again?"</div>
              </div>
            </div>
            <div>
              <h4 className="text-green-400 font-medium mb-2">✅ After gitr:</h4>
              <div className="space-y-2 text-sm text-gray-400">
                <div>• <code className="text-green-400">gitr clone &lt;url&gt;</code> → done</div>
                <div>• <code className="text-blue-400">Cmd+V</code> → navigate instantly</div>
                <div>• <code className="text-purple-400">gitr web</code> → browser opens</div>
                <div>• Always know exactly where everything is</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}