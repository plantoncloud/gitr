import React from 'react';
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Check, X, Terminal } from 'lucide-react';

const comparisons = [
  {
    feature: 'Deterministic clone paths',
    git: { status: 'none', text: 'Manual path selection every time' },
    gitr: { status: 'full', text: 'Automatic from URL structure' }
  },
  {
    feature: 'Multi-provider support',
    git: { status: 'manual', text: 'Same commands, different workflows' },
    gitr: { status: 'full', text: 'GitHub + GitLab + Bitbucket unified' }
  },
  {
    feature: 'Web navigation',
    git: { status: 'none', text: 'Copy/paste URLs manually' },
    gitr: { status: 'full', text: 'One command to any page' }
  },
  {
    feature: 'Fork handling',
    git: { status: 'manual', text: 'Path conflicts and confusion' },
    gitr: { status: 'full', text: 'Automatic path separation' }
  },
  {
    feature: 'Dry run preview',
    git: { status: 'none', text: 'No preview of clone location' },
    gitr: { status: 'full', text: 'See exactly what will happen' }
  },
  {
    feature: 'Clipboard integration',
    git: { status: 'none', text: 'Remember or type paths' },
    gitr: { status: 'full', text: 'Auto-copy cd commands' }
  }
];

const StatusIcon = ({ status }) => {
  switch (status) {
    case 'full':
      return <Check className="w-5 h-5 text-green-400" />;
    case 'manual':
      return <div className="w-5 h-5 rounded-full bg-yellow-500 flex items-center justify-center">
        <div className="w-2 h-2 bg-yellow-900 rounded-full" />
      </div>;
    case 'none':
      return <X className="w-5 h-5 text-red-400" />;
    default:
      return null;
  }
};

export default function CompareSection() {
  return (
    <div className="max-w-7xl mx-auto px-6">
      <div className="text-center mb-16">
        <h2 className="text-4xl md:text-5xl font-bold mb-6 bg-gradient-to-r from-white to-gray-300 bg-clip-text text-transparent">
          gitr vs git
        </h2>
        <p className="text-xl text-gray-400 max-w-3xl mx-auto">
          See how gitr enhances your standard git workflow
        </p>
      </div>

      <div className="max-w-4xl mx-auto">
        {/* Comparison Grid */}
        <div className="bg-gradient-to-r from-gray-900/50 to-gray-800/50 backdrop-blur-xl border border-gray-700 rounded-3xl overflow-hidden">
          {/* Header */}
          <div className="grid grid-cols-3 bg-gray-900/70 border-b border-gray-700">
            <div className="p-6">
              <h3 className="text-lg font-semibold text-white">Feature</h3>
            </div>
            <div className="p-6 border-l border-gray-700">
              <h3 className="text-lg font-semibold text-gray-400 flex items-center gap-2">
                <Terminal className="w-5 h-5" />
                Standard git
              </h3>
            </div>
            <div className="p-6 border-l border-gray-700">
              <h3 className="text-lg font-semibold text-blue-400 flex items-center gap-2">
                <Terminal className="w-5 h-5" />
                gitr
              </h3>
            </div>
          </div>

          {/* Comparison Rows */}
          {comparisons.map((item, index) => (
            <div 
              key={index} 
              className={`grid grid-cols-3 ${index % 2 === 0 ? 'bg-gray-900/30' : 'bg-gray-800/30'}`}
            >
              <div className="p-6">
                <span className="font-medium text-white">{item.feature}</span>
              </div>
              <div className="p-6 border-l border-gray-700">
                <div className="flex items-center gap-3">
                  <StatusIcon status={item.git.status} />
                  <span className="text-gray-400 text-sm">{item.git.text}</span>
                </div>
              </div>
              <div className="p-6 border-l border-gray-700">
                <div className="flex items-center gap-3">
                  <StatusIcon status={item.gitr.status} />
                  <span className="text-green-300 text-sm">{item.gitr.text}</span>
                </div>
              </div>
            </div>
          ))}
        </div>

        {/* Example Workflow */}
        <div className="mt-12 grid md:grid-cols-2 gap-8">
          <Card className="bg-gray-900/50 border-gray-700">
            <CardHeader>
              <h3 className="text-lg font-semibold text-white flex items-center gap-2">
                <Terminal className="w-5 h-5 text-red-400" />
                With standard git
              </h3>
            </CardHeader>
            <CardContent className="space-y-3">
              <div className="bg-black/50 rounded-lg p-3">
                <code className="text-red-400 text-sm">$ cd ~/where/should/this/go?</code>
              </div>
              <div className="bg-black/50 rounded-lg p-3">
                <code className="text-red-400 text-sm">$ git clone https://github.com/owner/repo</code>
              </div>
              <div className="bg-black/50 rounded-lg p-3">
                <code className="text-red-400 text-sm">$ cd repo</code>
              </div>
              <div className="bg-black/50 rounded-lg p-3">
                <code className="text-gray-500 text-sm"># Open browser manually...</code>
              </div>
            </CardContent>
          </Card>

          <Card className="bg-gray-900/50 border-gray-700">
            <CardHeader>
              <h3 className="text-lg font-semibold text-white flex items-center gap-2">
                <Terminal className="w-5 h-5 text-green-400" />
                With gitr
              </h3>
            </CardHeader>
            <CardContent className="space-y-3">
              <div className="bg-black/50 rounded-lg p-3">
                <code className="text-green-400 text-sm">$ gitr clone https://github.com/owner/repo</code>
              </div>
              <div className="bg-black/50 rounded-lg p-3">
                <code className="text-blue-400 text-sm">$ cmd+v  # paste copied path</code>
              </div>
              <div className="bg-black/50 rounded-lg p-3">
                <code className="text-purple-400 text-sm">$ gitr web  # opens in browser</code>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}