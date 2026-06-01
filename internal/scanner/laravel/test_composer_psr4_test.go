//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestComposerPSR4 테스트
package laravel

import "testing"

func TestComposerPSR4(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "composer.json", `{
		"autoload": { "psr-4": { "App\\": "app/", "Domain\\": ["src/Domain/", "x/"] } },
		"autoload-dev": { "psr-4": { "Tests\\": "tests/" } }
	}`)
	m := composerPSR4(dir)
	if m["App\\"] != "app/" {
		t.Errorf("App\\ = %q, want app/", m["App\\"])
	}
	if m["Domain\\"] != "src/Domain/" {
		t.Errorf("Domain\\ = %q, want src/Domain/", m["Domain\\"])
	}
	if m["Tests\\"] != "tests/" {
		t.Errorf("Tests\\ = %q, want tests/", m["Tests\\"])
	}
}
