package dto

type CveDTO struct {
	ResultsPerPage  int    `json:"resultsPerPage,omitempty"`
	StartIndex      int    `json:"startIndex,omitempty"`
	TotalResults    int    `json:"totalResults,omitempty"`
	Format          string `json:"format,omitempty"`
	Timestamp       string `json:"timestamp,omitempty"`
	Version         string `json:"version,omitempty"`
	Vulnerabilities []cve  `json:"vulnerabilities"`
}

type cve struct {
	Info struct {
		ID               string `json:"id,omitempty"`
		SourceIdentifier string `json:"sourceIdentifier,omitempty"`
		Published        string `json:"published,omitempty"`
		LastModified     string `json:"lastModified,omitempty"`
		VulnStatus       string `json:"vulnStatus,omitempty"`
		Descriptions     []struct {
			Lang  string `json:"lang,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"descriptions"`
		Metrics    metrics `json:"metrics"`
		Weaknesses []struct {
			Source      string `json:"source,omitempty"`
			Type        string `json:"type,omitempty"`
			Description []struct {
				Lang  string `json:"lang,omitempty" json:"lang,omitempty"`
				Value string `json:"value,omitempty" json:"value,omitempty"`
			} `json:"description,omitempty"`
		} `json:"weaknesses"`
		Configurations []struct {
			Operator string `json:"operator,omitempty"`
			Negate   bool   `json:"negate,omitempty"`
			CpeMatch []struct {
				Vulnerable      bool   `json:"vulnerable,omitempty"`
				Criteria        string `json:"criteria,omitempty"`
				MatchCriteriaId string `json:"matchCriteriaId,omitempty"`
			} `json:"cpeMatch,omitempty"`
		} `json:"configurations"`
		References []struct {
			Url    string `json:"url,omitempty"`
			Source string `json:"source,omitempty"`
		} `json:"references"`
	} `json:"cve"`
}

type metrics struct {
	Cm []struct {
		Source   string `json:"source,omitempty"`
		Type     string `json:"type,omitempty"`
		CvssData struct {
			Version               string      `json:"version,omitempty"`
			VectorString          string      `json:"vectorString,omitempty"`
			AccessVector          string      `json:"accessVector,omitempty"`
			AccessComplexity      string      `json:"accessComplexity,omitempty"`
			Authentication        string      `json:"authentication,omitempty"`
			ConfidentialityImpact string      `json:"confidentialityImpact,omitempty"`
			IntegrityImpact       string      `json:"integrityImpact,omitempty"`
			AvailabilityImpact    string      `json:"availabilityImpact,omitempty"`
			BaseScore             interface{} `json:"baseScore,omitempty"`
		} `json:"cvssData"`
		BaseSeverity            string      `json:"baseSeverity,omitempty"`
		ExploitabilityScore     interface{} `json:"exploitabilityScore,omitempty"`
		ImpactScore             interface{} `json:"impactScore,omitempty"`
		AcInsufInfo             bool        `json:"acInsufInfo,omitempty"`
		ObtainAllPrivilege      bool        `json:"obtainAllPrivilege,omitempty"`
		ObtainUserPrivilege     bool        `json:"obtainUserPrivilege,omitempty"`
		ObtainOtherPrivilege    bool        `json:"obtainOtherPrivilege,omitempty"`
		UserInteractionRequired bool        `json:"userInteractionRequired,omitempty"`
	} `json:"cvssMetricV2"`
}
