//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what writePSR4NonstandardProject — 비표준 PSR-4 위치의 컨트롤러/요청/리소스를 가진 테스트 프로젝트를 작성한다
package laravel

import "testing"

func writePSR4NonstandardProject(t *testing.T, dir string) {
	t.Helper()
	writeFile(t, dir, "composer.json", `{
		"autoload": { "psr-4": { "Acme\\Billing\\": "modules/billing/src/" } }
	}`)
	writeFile(t, dir, "routes/api.php", `<?php
use Acme\Billing\Controllers\InvoiceController;

Route::get('/invoices/{invoice}', [InvoiceController::class, 'show']);
Route::post('/invoices', [InvoiceController::class, 'store']);
`)
	writeFile(t, dir, "modules/billing/src/Controllers/InvoiceController.php", `<?php
namespace Acme\Billing\Controllers;

use Acme\Billing\Requests\StoreInvoiceRequest;
use Acme\Billing\Resources\InvoiceResource;

class InvoiceController {
    public function show(int $invoice) {
        return new InvoiceResource($x);
    }
    public function store(StoreInvoiceRequest $request) {
        return new InvoiceResource($x);
    }
}
`)
	writeFile(t, dir, "modules/billing/src/Requests/StoreInvoiceRequest.php", `<?php
namespace Acme\Billing\Requests;

use Illuminate\Foundation\Http\FormRequest;

class StoreInvoiceRequest extends FormRequest {
    public function rules(): array {
        return [
            'amount' => 'required|numeric',
            'currency' => 'required|string|max:3',
        ];
    }
}
`)
	writeFile(t, dir, "modules/billing/src/Resources/InvoiceResource.php", `<?php
namespace Acme\Billing\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class InvoiceResource extends JsonResource {
    public function toArray($request): array {
        return [
            'id' => $this->id,
            'amount' => $this->amount,
        ];
    }
}
`)
}
