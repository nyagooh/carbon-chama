document.addEventListener('DOMContentLoaded', function() {
  fetch('/api/chamas')
    .then(response => response.json())
    .then(data => {
      const chamaList = document.getElementById('chama-list');
      data.forEach(chama => {
        const card = document.createElement('div');
        card.className = 'bg-white rounded-lg shadow-md p-6';
        card.innerHTML = `
          <h2 class="text-xl font-semibold mb-2">${chama.name}</h2>
          <p class="text-gray-600 mb-4">${chama.description}</p>
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-500">Members: ${chama.members.length}</span>
            <button 
              onclick="joinChama('${chama.id}')"
              class="bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600"
            >
              Join
            </button>
          </div>
        `;
        chamaList.appendChild(card);
      });
    });
});

function joinChama(chamaId) {
  fetch('/api/chamas/join', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ chamaId })
  })
  .then(response => response.json())
  .then(data => {
    if (data.success) {
      alert('Successfully joined the chama!');
    } else {
      alert('Failed to join: ' + data.message);
    }
  });
}

async function saveForestChama() {
    if (typeof window.ethereum === 'undefined') {
        alert('Please install MetaMask!');
        return;
    }

    const accounts = await ethereum.request({ method: 'eth_requestAccounts' });
    const account = accounts[0];

    const contractAddress = '0xYourContractAddress';
    const abi = [/* Your ABI here */];

    const contract = new ethers.Contract(contractAddress, abi, new ethers.providers.Web3Provider(window.ethereum).getSigner());

    try {
        const tx = await contract.createContract(100); // 100 carbon credits
        await tx.wait();
        alert('Contract saved successfully!');
    } catch (error) {
        console.error(error);
        alert('Failed to save contract');
    }
}

// Add event listener to save button
document.addEventListener('DOMContentLoaded', function() {
  document.getElementById('save-forest-chama').addEventListener('click', saveForestChama);
});
